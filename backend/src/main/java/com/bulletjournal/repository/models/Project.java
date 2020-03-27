package com.bulletjournal.repository.models;

import com.bulletjournal.controller.models.ProjectType;
import org.hibernate.annotations.OnDelete;
import org.hibernate.annotations.OnDeleteAction;

import javax.persistence.*;
import javax.validation.constraints.NotNull;

@Entity
@Table(name = "projects",
        indexes = {@Index(name = "project_owner_index", columnList = "owner")},
        uniqueConstraints = {
                @UniqueConstraint(columnNames = {"owner", "name"})
        })
public class Project extends OwnedModel {

    @Id
    @GeneratedValue(generator = "project_generator")
    @SequenceGenerator(
            name = "project_generator",
            sequenceName = "project_sequence",
            initialValue = 100
    )
    private Long id;

    @NotNull
    @Column(updatable = false, nullable = false)
    private Integer type;

    @ManyToOne(fetch = FetchType.EAGER, optional = false)
    @JoinColumn(name = "group_id", nullable = false)
    @OnDelete(action = OnDeleteAction.NO_ACTION)
    private Group group;

    @Column
    private String description;

    @Column(nullable = false, columnDefinition = "Boolean default 'false'")
    private boolean shared;

    public Project() {
    }

    public Project(String name, Integer type, Group group, boolean shared) {
        this.type = type;
        this.group = group;
        this.shared = shared;
        this.setName(name);
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Integer getType() {
        return type;
    }

    public void setType(Integer type) {
        this.type = type;
    }

    public Group getGroup() {
        return group;
    }

    public void setGroup(Group group) {
        if (!group.getProjects().contains(group)) {
            this.group = group;
            group.addProject(this);
        }
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public boolean isShared() {
        return shared;
    }

    public void setShared(boolean shared) {
        this.shared = shared;
    }

    public com.bulletjournal.controller.models.Project toPresentationModel() {
        return new com.bulletjournal.controller.models.Project(
                this.getId(), this.getName(), this.getOwner(),
                ProjectType.getType(this.getType()),
                this.getGroup().toPresentationModel(),
                this.getDescription());
    }

    public com.bulletjournal.controller.models.Project toVerbosePresentationModel() {
        return new com.bulletjournal.controller.models.Project(
                this.getId(), this.getName(), this.getOwner(),
                ProjectType.getType(this.getType()),
                this.getGroup().toVerbosePresentationModel(),
                this.getDescription());
    }
}